
/**
 * Samples for Environments CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/
     * Environments_CreateOrUpdate.json
     */
    /**
     * Sample code: Environments_CreateOrUpdate.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void environmentsCreateOrUpdate(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.environments().define("public").withExistingWorkspace("contoso-resources", "contoso", "default")
            .create();
    }
}
