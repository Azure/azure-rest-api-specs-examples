
/**
 * Samples for ServerlessRuntimes CheckDependencies.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/
     * ServerlessRuntimes_CheckDependencies_MinimumSet_Gen.json
     */
    /**
     * Sample code: ServerlessRuntimes_CheckDependencies_Min.
     * 
     * @param manager Entry point to InformaticaDataManagementManager.
     */
    public static void serverlessRuntimesCheckDependenciesMin(
        com.azure.resourcemanager.informaticadatamanagement.InformaticaDataManagementManager manager) {
        manager.serverlessRuntimes().checkDependenciesWithResponse("rgopenapi", "_-", "_2_",
            com.azure.core.util.Context.NONE);
    }
}
