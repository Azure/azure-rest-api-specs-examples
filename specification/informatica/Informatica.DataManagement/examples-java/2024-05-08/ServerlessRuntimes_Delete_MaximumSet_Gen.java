
/**
 * Samples for ServerlessRuntimes Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/informatica/Informatica.DataManagement/examples/2024-05-08/ServerlessRuntimes_Delete_MaximumSet_Gen
     * .json
     */
    /**
     * Sample code: ServerlessRuntimes_Delete.
     * 
     * @param manager Entry point to InformaticaDataManagementManager.
     */
    public static void serverlessRuntimesDelete(
        com.azure.resourcemanager.informaticadatamanagement.InformaticaDataManagementManager manager) {
        manager.serverlessRuntimes().delete("rgopenapi", "orgName", "serverlessRuntimeName",
            com.azure.core.util.Context.NONE);
    }
}
