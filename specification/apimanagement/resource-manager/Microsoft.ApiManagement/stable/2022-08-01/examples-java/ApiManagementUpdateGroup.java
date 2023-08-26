import com.azure.resourcemanager.apimanagement.models.GroupContract;

/** Samples for Group Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUpdateGroup.json
     */
    /**
     * Sample code: ApiManagementUpdateGroup.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUpdateGroup(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        GroupContract resource =
            manager
                .groups()
                .getWithResponse("rg1", "apimService1", "tempgroup", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().withDisplayName("temp group").withIfMatch("*").apply();
    }
}
