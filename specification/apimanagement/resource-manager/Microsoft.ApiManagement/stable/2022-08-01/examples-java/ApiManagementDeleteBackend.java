/** Samples for Backend Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteBackend.json
     */
    /**
     * Sample code: ApiManagementDeleteBackend.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteBackend(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .backends()
            .deleteWithResponse("rg1", "apimService1", "sfbackend", "*", com.azure.core.util.Context.NONE);
    }
}
