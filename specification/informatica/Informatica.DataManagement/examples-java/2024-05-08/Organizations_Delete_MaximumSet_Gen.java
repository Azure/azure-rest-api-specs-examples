
/**
 * Samples for Organizations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/informatica/Informatica.DataManagement/examples/2024-05-08/Organizations_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organizations_Delete.
     * 
     * @param manager Entry point to InformaticaDataManagementManager.
     */
    public static void organizationsDelete(
        com.azure.resourcemanager.informaticadatamanagement.InformaticaDataManagementManager manager) {
        manager.organizations().delete("rgopenapi", "_", com.azure.core.util.Context.NONE);
    }
}
