
/**
 * Samples for Organizations GetAllServerlessRuntimes.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/informatica/Informatica.DataManagement/examples/2024-05-08/
     * Organizations_GetAllServerlessRuntimes_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organizations_GetAllServerlessRuntimes.
     * 
     * @param manager Entry point to InformaticaDataManagementManager.
     */
    public static void organizationsGetAllServerlessRuntimes(
        com.azure.resourcemanager.informaticadatamanagement.InformaticaDataManagementManager manager) {
        manager.organizations().getAllServerlessRuntimesWithResponse("rgopenapi", "t",
            com.azure.core.util.Context.NONE);
    }
}
