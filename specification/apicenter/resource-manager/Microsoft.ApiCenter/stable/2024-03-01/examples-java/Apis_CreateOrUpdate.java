
/**
 * Samples for Apis CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/Apis_CreateOrUpdate.json
     */
    /**
     * Sample code: Apis_CreateOrUpdate.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void apisCreateOrUpdate(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.apis().define("echo-api").withExistingWorkspace("contoso-resources", "contoso", "default").create();
    }
}
