/** Samples for ApiManagementSkus List. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListSku.json
     */
    /**
     * Sample code: Lists all available Resource SKUs.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void listsAllAvailableResourceSKUs(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiManagementSkus().list(com.azure.core.util.Context.NONE);
    }
}
