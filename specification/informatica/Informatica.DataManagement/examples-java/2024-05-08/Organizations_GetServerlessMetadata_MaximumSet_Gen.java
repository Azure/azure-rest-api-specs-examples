
/**
 * Samples for Organizations GetServerlessMetadata.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/informatica/Informatica.DataManagement/examples/2024-05-08/
     * Organizations_GetServerlessMetadata_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organizations_GetServerlessMetadata.
     * 
     * @param manager Entry point to InformaticaDataManagementManager.
     */
    public static void organizationsGetServerlessMetadata(
        com.azure.resourcemanager.informaticadatamanagement.InformaticaDataManagementManager manager) {
        manager.organizations().getServerlessMetadataWithResponse("rgopenapi", "3_UC",
            com.azure.core.util.Context.NONE);
    }
}
