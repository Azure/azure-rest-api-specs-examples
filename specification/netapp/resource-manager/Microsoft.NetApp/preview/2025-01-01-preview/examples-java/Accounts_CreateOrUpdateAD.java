
import com.azure.resourcemanager.netapp.models.ActiveDirectory;
import java.util.Arrays;

/**
 * Samples for Accounts CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/preview/2025-01-01-preview/examples/
     * Accounts_CreateOrUpdateAD.json
     */
    /**
     * Sample code: Accounts_CreateOrUpdateWithActiveDirectory.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void
        accountsCreateOrUpdateWithActiveDirectory(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.accounts().define("account1").withRegion("eastus").withExistingResourceGroup("myRG")
            .withActiveDirectories(Arrays.asList(new ActiveDirectory().withUsername("ad_user_name")
                .withPassword("fakeTokenPlaceholder").withDomain("10.10.10.3").withDns("10.10.10.3")
                .withSmbServerName("SMBServer").withOrganizationalUnit("OU=Engineering").withSite("SiteName")
                .withAesEncryption(true).withLdapSigning(false).withLdapOverTls(false)))
            .create();
    }
}
