/** Samples for Backend GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadBackend.json
     */
    /**
     * Sample code: ApiManagementHeadBackend.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadBackend(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .backends()
            .getEntityTagWithResponse("rg1", "apimService1", "sfbackend", com.azure.core.util.Context.NONE);
    }
}
