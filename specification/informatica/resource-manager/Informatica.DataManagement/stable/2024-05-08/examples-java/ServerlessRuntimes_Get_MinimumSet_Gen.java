
/**
 * Samples for ServerlessRuntimes Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/
     * ServerlessRuntimes_Get_MinimumSet_Gen.json
     */
    /**
     * Sample code: ServerlessRuntimes_Get_Min.
     * 
     * @param manager Entry point to InformaticaDataManagementManager.
     */
    public static void serverlessRuntimesGetMin(
        com.azure.resourcemanager.informaticadatamanagement.InformaticaDataManagementManager manager) {
        manager.serverlessRuntimes().getWithResponse("rgopenapi", "YC", "___", com.azure.core.util.Context.NONE);
    }
}
