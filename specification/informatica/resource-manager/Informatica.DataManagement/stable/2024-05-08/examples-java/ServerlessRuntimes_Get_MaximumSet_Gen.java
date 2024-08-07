
/**
 * Samples for ServerlessRuntimes Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/
     * ServerlessRuntimes_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: ServerlessRuntimes_Get.
     * 
     * @param manager Entry point to InformaticaDataManagementManager.
     */
    public static void serverlessRuntimesGet(
        com.azure.resourcemanager.informaticadatamanagement.InformaticaDataManagementManager manager) {
        manager.serverlessRuntimes().getWithResponse("rgopenapi", "e3Y", "48-", com.azure.core.util.Context.NONE);
    }
}
