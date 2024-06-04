
/**
 * Samples for ServerlessRuntimes StartFailedServerlessRuntime.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/informatica/Informatica.DataManagement/examples/2024-05-08/
     * ServerlessRuntimes_StartFailedServerlessRuntime_MaximumSet_Gen.json
     */
    /**
     * Sample code: ServerlessRuntimes_StartFailedServerlessRuntime.
     * 
     * @param manager Entry point to InformaticaDataManagementManager.
     */
    public static void serverlessRuntimesStartFailedServerlessRuntime(
        com.azure.resourcemanager.informaticadatamanagement.InformaticaDataManagementManager manager) {
        manager.serverlessRuntimes().startFailedServerlessRuntimeWithResponse("rgopenapi", "9M4", "-25-G_",
            com.azure.core.util.Context.NONE);
    }
}
