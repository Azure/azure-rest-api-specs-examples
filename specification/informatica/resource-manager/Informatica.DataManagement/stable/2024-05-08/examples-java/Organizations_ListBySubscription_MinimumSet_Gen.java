
/**
 * Samples for Organizations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/
     * Organizations_ListBySubscription_MinimumSet_Gen.json
     */
    /**
     * Sample code: Organizations_ListBySubscription_Min.
     * 
     * @param manager Entry point to InformaticaDataManagementManager.
     */
    public static void organizationsListBySubscriptionMin(
        com.azure.resourcemanager.informaticadatamanagement.InformaticaDataManagementManager manager) {
        manager.organizations().list(com.azure.core.util.Context.NONE);
    }
}
