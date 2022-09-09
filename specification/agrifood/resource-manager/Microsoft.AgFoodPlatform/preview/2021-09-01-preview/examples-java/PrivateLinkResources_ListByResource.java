import com.azure.core.util.Context;

/** Samples for PrivateLinkResources ListByResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/PrivateLinkResources_ListByResource.json
     */
    /**
     * Sample code: PrivateLinkResources_ListByResource.
     *
     * @param manager Entry point to AgriFoodManager.
     */
    public static void privateLinkResourcesListByResource(com.azure.resourcemanager.agrifood.AgriFoodManager manager) {
        manager.privateLinkResources().listByResource("examples-rg", "examples-farmbeatsResourceName", Context.NONE);
    }
}
