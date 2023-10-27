import com.azure.resourcemanager.applicationinsights.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.applicationinsights.models.UserAssignedIdentity;
import com.azure.resourcemanager.applicationinsights.models.WorkbookResourceIdentity;
import com.azure.resourcemanager.applicationinsights.models.WorkbookSharedTypeKind;
import java.util.HashMap;
import java.util.Map;

/** Samples for Workbooks CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2022-04-01/examples/WorkbookManagedAdd.json
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
            .withIdentity(
                new WorkbookResourceIdentity()
                    .withType(ManagedServiceIdentityType.USER_ASSIGNED)
                    .withUserAssignedIdentities(
                        mapOf(
                            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/my-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myid",
                            new UserAssignedIdentity())))
            .withKind(WorkbookSharedTypeKind.SHARED)
            .withDisplayName("Sample workbook")
            .withSerializedData(
                "{\"version\":\"Notebook/1.0\",\"items\":[{\"type\":1,\"content\":{\"json\":\"test\"},\"name\":\"text -"
                    + " 0\"}],\"isLocked\":false,\"fallbackResourceIds\":[\"/subscriptions/00000000-0000-0000-0000-00000000/resourceGroups/my-resource-group\"]}")
            .withVersion("Notebook/1.0")
            .withCategory("workbook")
            .withStorageUri(
                "/subscriptions/6b643656-33eb-422f-aee8-3ac145d124af/resourceGroups/my-resource-group/providers/Microsoft.Storage/storageAccounts/mystorage/blobServices/default/containers/mycontainer")
            .withDescription("Sample workbook")
            .withSourceIdParameter(
                "/subscriptions/6b643656-33eb-422f-aee8-3ac145d124af/resourcegroups/my-resource-group")
            .create();
    }

    // Use "Map.of" if available
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
