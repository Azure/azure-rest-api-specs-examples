
import com.azure.resourcemanager.sphere.models.ListDeviceGroupsRequest;

/**
 * Samples for Catalogs ListDeviceGroups.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/stable/2024-04-01/examples/
     * PostListDeviceGroupsCatalog.json
     */
    /**
     * Sample code: Catalogs_ListDeviceGroups.
     * 
     * @param manager Entry point to AzureSphereManager.
     */
    public static void catalogsListDeviceGroups(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager.catalogs().listDeviceGroups("MyResourceGroup1", "MyCatalog1",
            new ListDeviceGroupsRequest().withDeviceGroupName("MyDeviceGroup1"), null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
