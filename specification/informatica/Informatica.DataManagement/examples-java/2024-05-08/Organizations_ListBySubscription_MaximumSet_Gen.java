
/**
 * Samples for Organizations List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/informatica/Informatica.DataManagement/examples/2024-05-08/
     * Organizations_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organizations_ListBySubscription.
     * 
     * @param manager Entry point to InformaticaDataManagementManager.
     */
    public static void organizationsListBySubscription(
        com.azure.resourcemanager.informaticadatamanagement.InformaticaDataManagementManager manager) {
        manager.organizations().list(com.azure.core.util.Context.NONE);
    }
}
