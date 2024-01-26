
import com.azure.resourcemanager.appplatform.fluent.models.ApplicationAcceleratorResourceInner;
import com.azure.resourcemanager.appplatform.models.ApplicationAcceleratorProperties;
import com.azure.resourcemanager.appplatform.models.Sku;

/**
 * Samples for ApplicationAccelerators CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * ApplicationAccelerators_CreateOrUpdate.json
     */
    /**
     * Sample code: ApplicationAccelerators_CreateOrUpdate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void applicationAcceleratorsCreateOrUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getApplicationAccelerators()
            .createOrUpdate("myResourceGroup", "myservice", "default",
                new ApplicationAcceleratorResourceInner().withProperties(new ApplicationAcceleratorProperties())
                    .withSku(new Sku().withName("E0").withTier("Enterprise").withCapacity(2)),
                com.azure.core.util.Context.NONE);
    }
}
