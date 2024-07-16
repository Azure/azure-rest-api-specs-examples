
/**
 * Samples for Organizations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/
     * Organizations_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: Organizations_Delete_Min.
     * 
     * @param manager Entry point to InformaticaDataManagementManager.
     */
    public static void organizationsDeleteMin(
        com.azure.resourcemanager.informaticadatamanagement.InformaticaDataManagementManager manager) {
        manager.organizations().delete("rgopenapi", "_-", com.azure.core.util.Context.NONE);
    }
}
