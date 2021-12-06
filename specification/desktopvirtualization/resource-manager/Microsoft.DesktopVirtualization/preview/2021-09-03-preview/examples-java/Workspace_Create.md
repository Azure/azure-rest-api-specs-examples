Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-desktopvirtualization_1.0.0-beta.1/sdk/desktopvirtualization/azure-resourcemanager-desktopvirtualization/README.md) on how to add the SDK to your project and authenticate.

```java
import java.util.HashMap;
import java.util.Map;

/** Samples for Workspaces CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/Workspace_Create.json
     */
    /**
     * Sample code: Workspace_Create.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void workspaceCreate(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .workspaces()
            .define("workspace1")
            .withRegion("centralus")
            .withExistingResourceGroup("resourceGroup1")
            .withTags(mapOf("tag1", "value1", "tag2", "value2"))
            .withDescription("des1")
            .withFriendlyName("friendly")
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
