
/**
 * Samples for ServerlessRuntimes CheckDependencies.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/
     * ServerlessRuntimes_CheckDependencies_MaximumSet_Gen.json
     */
    /**
     * Sample code: ServerlessRuntimes_CheckDependencies.
     * 
     * @param manager Entry point to InformaticaDataManagementManager.
     */
    public static void serverlessRuntimesCheckDependencies(
        com.azure.resourcemanager.informaticadatamanagement.InformaticaDataManagementManager manager) {
        manager.serverlessRuntimes().checkDependenciesWithResponse("rgopenapi", "3P", "M",
            com.azure.core.util.Context.NONE);
    }
}
