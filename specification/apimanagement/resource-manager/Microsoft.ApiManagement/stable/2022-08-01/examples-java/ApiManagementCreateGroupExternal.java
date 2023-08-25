import com.azure.resourcemanager.apimanagement.models.GroupType;

/** Samples for Group CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateGroupExternal.json
     */
    /**
     * Sample code: ApiManagementCreateGroupExternal.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateGroupExternal(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .groups()
            .define("aadGroup")
            .withExistingService("rg1", "apimService1")
            .withDisplayName("NewGroup (samiraad.onmicrosoft.com)")
            .withDescription("new group to test")
            .withType(GroupType.EXTERNAL)
            .withExternalId("aad://samiraad.onmicrosoft.com/groups/83cf2753-5831-4675-bc0e-2f8dc067c58d")
            .create();
    }
}
