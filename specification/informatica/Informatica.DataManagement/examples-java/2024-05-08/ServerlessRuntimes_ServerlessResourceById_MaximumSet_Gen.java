
/**
 * Samples for ServerlessRuntimes ServerlessResourceById.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/informatica/Informatica.DataManagement/examples/2024-05-08/
     * ServerlessRuntimes_ServerlessResourceById_MaximumSet_Gen.json
     */
    /**
     * Sample code: ServerlessRuntimes_ServerlessResourceById.
     * 
     * @param manager Entry point to InformaticaDataManagementManager.
     */
    public static void serverlessRuntimesServerlessResourceById(
        com.azure.resourcemanager.informaticadatamanagement.InformaticaDataManagementManager manager) {
        manager.serverlessRuntimes().serverlessResourceByIdWithResponse("rgopenapi", "_RD_R",
            "serverlessRuntimeName159", com.azure.core.util.Context.NONE);
    }
}
