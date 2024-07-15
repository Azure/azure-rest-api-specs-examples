
/**
 * Samples for ServerlessRuntimes CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/
     * ServerlessRuntimes_CreateOrUpdate_MinimumSet_Gen.json
     */
    /**
     * Sample code: ServerlessRuntimes_CreateOrUpdate_Min.
     * 
     * @param manager Entry point to InformaticaDataManagementManager.
     */
    public static void serverlessRuntimesCreateOrUpdateMin(
        com.azure.resourcemanager.informaticadatamanagement.InformaticaDataManagementManager manager) {
        manager.serverlessRuntimes().define("J").withExistingOrganization("rgopenapi", "-4Z__7").create();
    }
}
