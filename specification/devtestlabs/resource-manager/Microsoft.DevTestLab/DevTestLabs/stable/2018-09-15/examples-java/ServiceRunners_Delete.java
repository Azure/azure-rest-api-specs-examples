
/**
 * Samples for ServiceRunners Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/ServiceRunners_Delete.json
     */
    /**
     * Sample code: ServiceRunners_Delete.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void serviceRunnersDelete(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.serviceRunners().deleteWithResponse("resourceGroupName", "{devtestlabName}", "{servicerunnerName}",
            com.azure.core.util.Context.NONE);
    }
}
