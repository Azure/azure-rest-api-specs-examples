/** Samples for Catalogs ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/GetCatalogsRG.json
     */
    /**
     * Sample code: Catalogs_ListByResourceGroup.
     *
     * @param manager Entry point to AzureSphereManager.
     */
    public static void catalogsListByResourceGroup(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager.catalogs().listByResourceGroup("MyResourceGroup1", com.azure.core.util.Context.NONE);
    }
}
