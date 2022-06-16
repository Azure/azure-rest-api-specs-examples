import com.azure.core.util.Context;
import com.azure.resourcemanager.imagebuilder.models.ImageTemplate;
import com.azure.resourcemanager.imagebuilder.models.ImageTemplateIdentity;
import com.azure.resourcemanager.imagebuilder.models.ResourceIdentityType;
import java.util.HashMap;
import java.util.Map;

/** Samples for VirtualMachineImageTemplates Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2022-02-14/examples/UpdateImageTemplateToRemoveIdentities.json
     */
    /**
     * Sample code: Remove identities for an Image Template.
     *
     * @param manager Entry point to ImageBuilderManager.
     */
    public static void removeIdentitiesForAnImageTemplate(
        com.azure.resourcemanager.imagebuilder.ImageBuilderManager manager) {
        ImageTemplate resource =
            manager
                .virtualMachineImageTemplates()
                .getByResourceGroupWithResponse("myResourceGroup", "myImageTemplate", Context.NONE)
                .getValue();
        resource.update().withIdentity(new ImageTemplateIdentity().withType(ResourceIdentityType.NONE)).apply();
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
