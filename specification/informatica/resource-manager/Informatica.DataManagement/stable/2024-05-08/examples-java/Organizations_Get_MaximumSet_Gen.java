
/**
 * Samples for Organizations GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/
     * Organizations_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organizations_Get.
     * 
     * @param manager Entry point to InformaticaDataManagementManager.
     */
    public static void
        organizationsGet(com.azure.resourcemanager.informaticadatamanagement.InformaticaDataManagementManager manager) {
        manager.organizations().getByResourceGroupWithResponse("rgopenapi", "Sg", com.azure.core.util.Context.NONE);
    }
}
