import com.azure.core.util.Context;

/** Samples for Backend Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetBackend.json
     */
    /**
     * Sample code: ApiManagementGetBackend.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetBackend(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.backends().getWithResponse("rg1", "apimService1", "sfbackend", Context.NONE);
    }
}
