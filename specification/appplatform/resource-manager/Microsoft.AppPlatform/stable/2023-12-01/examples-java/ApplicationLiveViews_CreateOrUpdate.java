
import com.azure.resourcemanager.appplatform.fluent.models.ApplicationLiveViewResourceInner;
import com.azure.resourcemanager.appplatform.models.ApplicationLiveViewProperties;

/**
 * Samples for ApplicationLiveViews CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * ApplicationLiveViews_CreateOrUpdate.json
     */
    /**
     * Sample code: ApplicationLiveViews_CreateOrUpdate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void applicationLiveViewsCreateOrUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getApplicationLiveViews().createOrUpdate("myResourceGroup",
            "myservice", "default",
            new ApplicationLiveViewResourceInner().withProperties(new ApplicationLiveViewProperties()),
            com.azure.core.util.Context.NONE);
    }
}
