
/**
 * Samples for Organizations GetServerlessMetadata.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/
     * Organizations_GetServerlessMetadata_MinimumSet_Gen.json
     */
    /**
     * Sample code: Organizations_GetServerlessMetadata_Min.
     * 
     * @param manager Entry point to InformaticaDataManagementManager.
     */
    public static void organizationsGetServerlessMetadataMin(
        com.azure.resourcemanager.informaticadatamanagement.InformaticaDataManagementManager manager) {
        manager.organizations().getServerlessMetadataWithResponse("rgopenapi", "A", com.azure.core.util.Context.NONE);
    }
}
