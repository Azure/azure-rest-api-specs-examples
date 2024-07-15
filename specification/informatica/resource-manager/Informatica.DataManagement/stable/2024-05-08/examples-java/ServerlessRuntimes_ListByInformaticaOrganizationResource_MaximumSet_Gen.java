
/**
 * Samples for ServerlessRuntimes ListByInformaticaOrganizationResource.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/
     * ServerlessRuntimes_ListByInformaticaOrganizationResource_MaximumSet_Gen.json
     */
    /**
     * Sample code: ServerlessRuntimes_ListByInformaticaOrganizationResource.
     * 
     * @param manager Entry point to InformaticaDataManagementManager.
     */
    public static void serverlessRuntimesListByInformaticaOrganizationResource(
        com.azure.resourcemanager.informaticadatamanagement.InformaticaDataManagementManager manager) {
        manager.serverlessRuntimes().listByInformaticaOrganizationResource("rgopenapi", "orgName",
            com.azure.core.util.Context.NONE);
    }
}
