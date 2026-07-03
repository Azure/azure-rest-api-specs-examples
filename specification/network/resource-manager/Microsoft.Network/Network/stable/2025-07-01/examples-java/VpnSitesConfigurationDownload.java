
import com.azure.resourcemanager.network.models.GetVpnSitesConfigurationRequest;
import java.util.Arrays;

/**
 * Samples for VpnSitesConfiguration Download.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VpnSitesConfigurationDownload.json
     */
    /**
     * Sample code: VpnSitesConfigurationDownload.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void vpnSitesConfigurationDownload(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVpnSitesConfigurations().download("rg1", "wan1",
            new GetVpnSitesConfigurationRequest().withVpnSites(Arrays.asList(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/vpnSites/abc"))
                .withOutputBlobSasUrl(
                    "https://blobcortextesturl.blob.core.windows.net/folderforconfig/vpnFile?sp=rw&se=2018-01-10T03%3A42%3A04Z&sv=2017-04-17&sig=WvXrT5bDmDFfgHs%2Brz%2BjAu123eRCNE9BO0eQYcPDT7pY%3D&sr=b"),
            com.azure.core.util.Context.NONE);
    }
}
