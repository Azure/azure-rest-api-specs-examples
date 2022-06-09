```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.imagebuilder.models.ImageTemplate;
import java.util.HashMap;
import java.util.Map;

/** Samples for VirtualMachineImageTemplates Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2022-02-14/examples/UpdateImageTemplateTags.json
     */
    /**
     * Sample code: Update the tags for an Image Template.
     *
     * @param manager Entry point to ImageBuilderManager.
     */
    public static void updateTheTagsForAnImageTemplate(
        com.azure.resourcemanager.imagebuilder.ImageBuilderManager manager) {
        ImageTemplate resource =
            manager
                .virtualMachineImageTemplates()
                .getByResourceGroupWithResponse("myResourceGroup", "myImageTemplate", Context.NONE)
                .getValue();
        resource.update().withTags(mapOf("new-tag", "new-value")).apply();
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
