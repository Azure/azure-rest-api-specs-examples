
import com.azure.resourcemanager.search.fluent.models.SharedPrivateLinkResourceInner;
import com.azure.resourcemanager.search.models.SharedPrivateLinkResourceProperties;

/**
 * Samples for SharedPrivateLinkResources CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/CreateOrUpdateSharedPrivateLinkResource.json
     */
    /**
     * Sample code: SharedPrivateLinkResourceCreateOrUpdate.
     * 
     * @param manager Entry point to SearchServiceManager.
     */
    public static void
        sharedPrivateLinkResourceCreateOrUpdate(com.azure.resourcemanager.search.SearchServiceManager manager) {
        manager.serviceClient().getSharedPrivateLinkResources().createOrUpdate("rg1", "mysearchservice", "testResource",
            new SharedPrivateLinkResourceInner().withProperties(new SharedPrivateLinkResourceProperties()
                .withPrivateLinkResourceId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Storage/storageAccounts/storageAccountName")
                .withGroupId("blob").withRequestMessage("please approve")),
            com.azure.core.util.Context.NONE);
    }
}
