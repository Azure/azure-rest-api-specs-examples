Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-desktopvirtualization_1.0.0-beta.1/sdk/desktopvirtualization/azure-resourcemanager-desktopvirtualization/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.desktopvirtualization.models.ApplicationGroupType;
import com.azure.resourcemanager.desktopvirtualization.models.MigrationRequestProperties;
import com.azure.resourcemanager.desktopvirtualization.models.Operation;
import java.util.HashMap;
import java.util.Map;

/** Samples for ApplicationGroups CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/ApplicationGroup_Create.json
     */
    /**
     * Sample code: ApplicationGroup_Create.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void applicationGroupCreate(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .applicationGroups()
            .define("applicationGroup1")
            .withRegion("centralus")
            .withExistingResourceGroup("resourceGroup1")
            .withHostPoolArmPath(
                "/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/hostPools/hostPool1")
            .withApplicationGroupType(ApplicationGroupType.REMOTE_APP)
            .withTags(mapOf("tag1", "value1", "tag2", "value2"))
            .withDescription("des1")
            .withFriendlyName("friendly")
            .withMigrationRequest(
                new MigrationRequestProperties()
                    .withOperation(Operation.START)
                    .withMigrationPath(
                        "TenantGroups/{defaultV1TenantGroup.Name}/Tenants/{defaultV1Tenant.Name}/HostPools/{sessionHostPool.Name}"))
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
