Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-applicationinsights_1.0.0-beta.3/sdk/applicationinsights/azure-resourcemanager-applicationinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.applicationinsights.models.Kind;
import com.azure.resourcemanager.applicationinsights.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.applicationinsights.models.UserAssignedIdentity;
import com.azure.resourcemanager.applicationinsights.models.WorkbookResourceIdentity;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for Workbooks CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2021-08-01/examples/WorkbookManagedAdd.json
     */
    /**
     * Sample code: WorkbookManagedAdd.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void workbookManagedAdd(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .workbooks()
            .define("deadb33f-5e0d-4064-8ebb-1a4ed0313eb2")
            .withRegion("westus")
            .withExistingResourceGroup("my-resource-group")
            .withTags(mapOf("hidden-title", "tttt"))
            .withIdentity(
                new WorkbookResourceIdentity()
                    .withType(ManagedServiceIdentityType.USER_ASSIGNED)
                    .withUserAssignedIdentities(
                        mapOf(
                            "/subscriptions/6b643656-33eb-422f-aee8-3ac145d124af/resourceGroups/eu2cgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1",
                            new UserAssignedIdentity())))
            .withKind(Kind.SHARED)
            .withEtag("\"4a00f78d-0000-0700-0000-5f8f616c1000\"")
            .withDisplayName("tttt")
            .withSerializedData(
                "{\"version\":\"Notebook/1.0\",\"items\":[{\"type\":1,\"content\":{\"json\":\"test\"},\"name\":\"text -"
                    + " 0\"}],\"isLocked\":false,\"fallbackResourceIds\":[\"/subscriptions/00000000-0000-0000-0000-00000000/resourceGroups/MyGroup\"]}")
            .withVersion("Notebook/1.0")
            .withCategory("workbook")
            .withTagsPropertiesTags(Arrays.asList())
            .withStorageUri(
                "/subscriptions/6b643656-33eb-422f-aee8-3ac145d124af/resourceGroups/MyGroup/providers/Microsoft.Storage/storageAccounts/testStorage/blobServices/default/containers/testContainer")
            .withDescription("Sample workbook")
            .withSourceIdParameter("/subscriptions/6b643656-33eb-422f-aee8-3ac145d124af/resourcegroups/MyGroup")
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
