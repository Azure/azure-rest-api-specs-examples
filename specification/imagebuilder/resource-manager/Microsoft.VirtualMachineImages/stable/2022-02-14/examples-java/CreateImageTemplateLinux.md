```java
import com.azure.resourcemanager.imagebuilder.models.ImageTemplateIdentity;
import com.azure.resourcemanager.imagebuilder.models.ImageTemplateIdentityUserAssignedIdentities;
import com.azure.resourcemanager.imagebuilder.models.ImageTemplateManagedImageDistributor;
import com.azure.resourcemanager.imagebuilder.models.ImageTemplateManagedImageSource;
import com.azure.resourcemanager.imagebuilder.models.ImageTemplateShellCustomizer;
import com.azure.resourcemanager.imagebuilder.models.ImageTemplateVmProfile;
import com.azure.resourcemanager.imagebuilder.models.ResourceIdentityType;
import com.azure.resourcemanager.imagebuilder.models.VirtualNetworkConfig;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for VirtualMachineImageTemplates CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2022-02-14/examples/CreateImageTemplateLinux.json
     */
    /**
     * Sample code: Create an Image Template for Linux.
     *
     * @param manager Entry point to ImageBuilderManager.
     */
    public static void createAnImageTemplateForLinux(
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
                        new ImageTemplateShellCustomizer()
                            .withName("Shell Customizer Example")
                            .withScriptUri("https://example.com/path/to/script.sh")))
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-imagebuilder_1.0.0-beta.3/sdk/imagebuilder/azure-resourcemanager-imagebuilder/README.md) on how to add the SDK to your project and authenticate.
