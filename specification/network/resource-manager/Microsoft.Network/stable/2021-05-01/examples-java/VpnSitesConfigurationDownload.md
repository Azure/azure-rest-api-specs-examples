Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.models.GetVpnSitesConfigurationRequest;
import java.util.Arrays;

/** Samples for VpnSitesConfiguration Download. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VpnSitesConfigurationDownload.json
     */
    /**
     * Sample code: VpnSitesConfigurationDownload.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void vpnSitesConfigurationDownload(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVpnSitesConfigurations()
            .download(
                "rg1",
                "wan1",
                new GetVpnSitesConfigurationRequest()
                    .withVpnSites(
                        Arrays
                            .asList("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnSites/abc"))
                    .withOutputBlobSasUrl(
                        "https://blobcortextesturl.blob.core.windows.net/folderforconfig/vpnFile?sp=rw&se=2018-01-10T03%3A42%3A04Z&sv=2017-04-17&sig=WvXrT5bDmDFfgHs%2Brz%2BjAu123eRCNE9BO0eQYcPDT7pY%3D&sr=b"),
                Context.NONE);
    }
}
```
