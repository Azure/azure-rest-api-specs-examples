
/**
 * Samples for Organizations ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/informatica/Informatica.DataManagement/examples/2024-05-08/
     * Organizations_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organizations_ListByResourceGroup.
     * 
     * @param manager Entry point to InformaticaDataManagementManager.
     */
    public static void organizationsListByResourceGroup(
        com.azure.resourcemanager.informaticadatamanagement.InformaticaDataManagementManager manager) {
        manager.organizations().listByResourceGroup("rgopenapi", com.azure.core.util.Context.NONE);
    }
}
