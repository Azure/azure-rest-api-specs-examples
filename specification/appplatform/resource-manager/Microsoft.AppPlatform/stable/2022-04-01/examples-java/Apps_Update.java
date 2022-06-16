import com.azure.core.util.Context;
import com.azure.resourcemanager.appplatform.fluent.models.AppResourceInner;
import com.azure.resourcemanager.appplatform.models.AppResourceProperties;
import com.azure.resourcemanager.appplatform.models.ManagedIdentityProperties;
import com.azure.resourcemanager.appplatform.models.ManagedIdentityType;
import com.azure.resourcemanager.appplatform.models.PersistentDisk;
import com.azure.resourcemanager.appplatform.models.TemporaryDisk;

/** Samples for Apps Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/Apps_Update.json
     */
    /**
     * Sample code: Apps_Update.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void appsUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getApps()
            .update(
                "myResourceGroup",
                "myservice",
                "myapp",
                new AppResourceInner()
                    .withProperties(
                        new AppResourceProperties()
                            .withPublicProperty(true)
                            .withFqdn("myapp.mydomain.com")
                            .withHttpsOnly(false)
                            .withTemporaryDisk(new TemporaryDisk().withSizeInGB(2).withMountPath("/mytemporarydisk"))
                            .withPersistentDisk(new PersistentDisk().withSizeInGB(2).withMountPath("/mypersistentdisk"))
                            .withEnableEndToEndTls(false))
                    .withIdentity(new ManagedIdentityProperties().withType(ManagedIdentityType.SYSTEM_ASSIGNED))
                    .withLocation("eastus"),
                Context.NONE);
    }
}
