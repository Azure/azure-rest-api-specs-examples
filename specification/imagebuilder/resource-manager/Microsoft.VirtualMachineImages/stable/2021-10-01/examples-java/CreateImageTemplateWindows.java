import com.azure.resourcemanager.imagebuilder.models.ImageTemplateIdentity;
import com.azure.resourcemanager.imagebuilder.models.ImageTemplateIdentityUserAssignedIdentities;
import com.azure.resourcemanager.imagebuilder.models.ImageTemplateManagedImageDistributor;
import com.azure.resourcemanager.imagebuilder.models.ImageTemplateManagedImageSource;
import com.azure.resourcemanager.imagebuilder.models.ImageTemplatePowerShellCustomizer;
import com.azure.resourcemanager.imagebuilder.models.ImageTemplateRestartCustomizer;
import com.azure.resourcemanager.imagebuilder.models.ImageTemplateVmProfile;
import com.azure.resourcemanager.imagebuilder.models.ImageTemplateWindowsUpdateCustomizer;
import com.azure.resourcemanager.imagebuilder.models.ResourceIdentityType;
import com.azure.resourcemanager.imagebuilder.models.VirtualNetworkConfig;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for VirtualMachineImageTemplates CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2021-10-01/examples/CreateImageTemplateWindows.json
     */
    /**
     * Sample code: Create an Image Template for Windows.
     *
     * @param manager Entry point to ImageBuilderManager.
     */
    public static void createAnImageTemplateForWindows(
        com.azure.resourcemanager.imagebuilder.ImageBuilderManager manager) {
        manager
            .virtualMachineImageTemplates()
            .define("myImageTemplate")
            .withRegion("westus")
            .withExistingResourceGroup("myResourceGroup")
            .withIdentity(
                new ImageTemplateIdentity()
                    .withType(ResourceIdentityType.USER_ASSIGNED)
                    .withUserAssignedIdentities(
                        mapOf(
                            "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity_1",
                            new ImageTemplateIdentityUserAssignedIdentities())))
            .withTags(mapOf("imagetemplate_tag1", "IT_T1", "imagetemplate_tag2", "IT_T2"))
            .withSource(
                new ImageTemplateManagedImageSource()
                    .withImageId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/images/source_image"))
            .withCustomize(
                Arrays
                    .asList(
                        new ImageTemplatePowerShellCustomizer()
                            .withName("PowerShell (inline) Customizer Example")
                            .withInline(
                                Arrays.asList("Powershell command-1", "Powershell command-2", "Powershell command-3")),
                        new ImageTemplatePowerShellCustomizer()
                            .withName("PowerShell (inline) Customizer Elevated user Example")
                            .withInline(
                                Arrays.asList("Powershell command-1", "Powershell command-2", "Powershell command-3"))
                            .withRunElevated(true),
                        new ImageTemplatePowerShellCustomizer()
                            .withName("PowerShell (inline) Customizer Elevated Local System user Example")
                            .withInline(
                                Arrays.asList("Powershell command-1", "Powershell command-2", "Powershell command-3"))
                            .withRunElevated(true)
                            .withRunAsSystem(true),
                        new ImageTemplatePowerShellCustomizer()
                            .withName("PowerShell (script) Customizer Example")
                            .withScriptUri("https://example.com/path/to/script.ps1")
                            .withValidExitCodes(Arrays.asList(0, 1)),
                        new ImageTemplatePowerShellCustomizer()
                            .withName("PowerShell (script) Customizer Elevated Local System user Example")
                            .withScriptUri("https://example.com/path/to/script.ps1")
                            .withRunElevated(true)
                            .withValidExitCodes(Arrays.asList(0, 1)),
                        new ImageTemplatePowerShellCustomizer()
                            .withName("PowerShell (script) Customizer Elevated Local System user Example")
                            .withScriptUri("https://example.com/path/to/script.ps1")
                            .withRunElevated(true)
                            .withRunAsSystem(true)
                            .withValidExitCodes(Arrays.asList(0, 1)),
                        new ImageTemplateRestartCustomizer()
                            .withName("Restart Customizer Example")
                            .withRestartCommand("shutdown /f /r /t 0 /c \"packer restart\"")
                            .withRestartCheckCommand("powershell -command \"& {Write-Output 'restarted.'}\"")
                            .withRestartTimeout("10m"),
                        new ImageTemplateWindowsUpdateCustomizer()
                            .withName("Windows Update Customizer Example")
                            .withSearchCriteria("BrowseOnly=0 and IsInstalled=0")
                            .withFilters(Arrays.asList("$_.BrowseOnly"))
                            .withUpdateLimit(100)))
            .withDistribute(
                Arrays
                    .asList(
                        new ImageTemplateManagedImageDistributor()
                            .withRunOutputName("image_it_pir_1")
                            .withArtifactTags(mapOf("tagName", "value"))
                            .withImageId(
                                "/subscriptions/{subscription-id}/resourceGroups/rg1/providers/Microsoft.Compute/images/image_it_1")
                            .withLocation("1_location")))
            .withVmProfile(
                new ImageTemplateVmProfile()
                    .withVmSize("Standard_D2s_v3")
                    .withOsDiskSizeGB(64)
                    .withVnetConfig(
                        new VirtualNetworkConfig()
                            .withSubnetId(
                                "/subscriptions/{subscription-id}/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet_name/subnets/subnet_name")))
            .create();
    }

    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
