
/**
 * Samples for Organizations GetAllServerlessRuntimes.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/
     * Organizations_GetAllServerlessRuntimes_MinimumSet_Gen.json
     */
    /**
     * Sample code: Organizations_GetAllServerlessRuntimes_Min.
     * 
     * @param manager Entry point to InformaticaDataManagementManager.
     */
    public static void organizationsGetAllServerlessRuntimesMin(
        com.azure.resourcemanager.informaticadatamanagement.InformaticaDataManagementManager manager) {
        manager.organizations().getAllServerlessRuntimesWithResponse("rgopenapi", "0",
            com.azure.core.util.Context.NONE);
    }
}
