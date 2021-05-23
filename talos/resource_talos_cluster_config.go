package talos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/talos-systems/talos/cmd/talosctl/cmd/mgmt"
	"github.com/talos-systems/talos/pkg/machinery/config"
	"github.com/talos-systems/talos/pkg/machinery/config/encoder"
	"github.com/talos-systems/talos/pkg/machinery/config/types/v1alpha1/generate"
	"gopkg.in/yaml.v3"
)

func resourceTalosClusterConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceTalosClusterConfigCreate,
		Read:   resourceTalosClusterConfigRead,
		Delete: resourceTalosClusterConfigDelete,

		Schema: map[string]*schema.Schema{
			"cluster_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"endpoint": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"additional_sans": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: false,
				Optional: true,
				ForceNew: true,
			},
			"dns_domain": {
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Default:  "cluster.local",
				ForceNew: true,
			},
			"install_disk": {
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Default:  "/dev/sda",
				ForceNew: true,
			},
			"install_image": {
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Default:  "ghcr.io/talos-systems/installer:v0.10.1",
				ForceNew: true,
			},
			"kubernetes_version": {
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Default:  "",
				ForceNew: true,
			},
			"persist_config": {
				Type:     schema.TypeBool,
				Required: false,
				Optional: true,
				Default:  true,
				ForceNew: true,
			},
			"registry_mirrors": {
				Type:     schema.TypeMap,
				Required: false,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Default:  map[string]string{},
				ForceNew: true,
			},
			"talos_version": {
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Default:  "",
				ForceNew: true,
			},
			"config_patch": {
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Default:  "",
				ForceNew: true,
			},
			"config_patch_control_plane": {
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Default:  "",
				ForceNew: true,
			},
			"config_patch_join": {
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Default:  "",
				ForceNew: true,
			},
			"bootstrap_user_data": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"controlplane_user_data": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"join_user_data": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"talos_config": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceTalosClusterConfigCreate(d *schema.ResourceData, meta interface{}) error {
	clusterName := d.Get("cluster_name").(string)
	endpoint := d.Get("endpoint").(string)
	additionalSANs, ok := d.Get("additional_sans").([]string)
	if !ok {
		additionalSANs = []string{}
	}
	dnsDomain := d.Get("dns_domain").(string)
	installDisk := d.Get("install_disk").(string)
	installImage := d.Get("install_image").(string)
	kubernetesVersion := d.Get("kubernetes_version").(string)
	persistConfig := d.Get("persist_config").(bool)
	registryMirrors, ok := d.Get("registry_mirrors").(map[string]string)
	if !ok {
		registryMirrors = map[string]string{}
	}
	talosVersion := d.Get("talos_version").(string)
	configPatch := d.Get("config_patch").(string)
	configPatchControlPlane := d.Get("config_patch_control_plane").(string)
	configPatchJoin := d.Get("config_patch_join").(string)

	var options []generate.GenOption

	for registryHost, mirrorUrl := range registryMirrors {
		options = append(options, generate.WithRegistryMirror(registryHost, mirrorUrl))
	}

	if talosVersion != "" {
		versionContract, err := config.ParseContractFromVersion(talosVersion)
		if err != nil {
			return err
		}

		options = append(options, generate.WithVersionContract(versionContract))
	}

	options = []generate.GenOption{
		generate.WithInstallDisk(installDisk),
		generate.WithInstallImage(installImage),
		generate.WithAdditionalSubjectAltNames(additionalSANs),
		generate.WithDNSDomain(dnsDomain),
		generate.WithPersist(persistConfig),
	}

	configBundle, err := mgmt.GenV1Alpha1Config(options, clusterName, endpoint, kubernetesVersion, configPatch, configPatchControlPlane, configPatchJoin)
	if err != nil {
		return err
	}

	d.SetId(clusterName)

	encoderOptions := []encoder.Option{
		encoder.WithComments(encoder.CommentsDisabled),
	}

	bootstrapUserData, err := configBundle.Init().String(encoderOptions...)
	if err != nil {
		return err
	}
	controlPlaneUserData, err := configBundle.ControlPlane().String(encoderOptions...)
	if err != nil {
		return err
	}
	joinUserData, err := configBundle.Join().String(encoderOptions...)
	if err != nil {
		return err
	}
	talosConfigBytes, err := yaml.Marshal(configBundle.TalosConfig())
	if err != nil {
		return err
	}

	d.Set("bootstrap_user_data", bootstrapUserData)
	d.Set("controlplane_user_data", controlPlaneUserData)
	d.Set("join_user_data", joinUserData)
	d.Set("talos_config", string(talosConfigBytes))

	return nil
}

func resourceTalosClusterConfigRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceTalosClusterConfigDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
