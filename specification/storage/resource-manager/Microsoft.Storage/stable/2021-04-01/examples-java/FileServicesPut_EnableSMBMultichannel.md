Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.storage.fluent.models.FileServicePropertiesInner;
import com.azure.resourcemanager.storage.models.Multichannel;
import com.azure.resourcemanager.storage.models.ProtocolSettings;
import com.azure.resourcemanager.storage.models.SmbSetting;

/** Samples for FileServices SetServiceProperties. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-04-01/examples/FileServicesPut_EnableSMBMultichannel.json
     */
    /**
     * Sample code: PutFileServices_EnableSMBMultichannel.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void putFileServicesEnableSMBMultichannel(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getFileServices()
            .setServicePropertiesWithResponse(
                "res4410",
                "sto8607",
                new FileServicePropertiesInner()
                    .withProtocolSettings(
                        new ProtocolSettings()
                            .withSmb(new SmbSetting().withMultichannel(new Multichannel().withEnabled(true)))),
                Context.NONE);
    }
}
```
