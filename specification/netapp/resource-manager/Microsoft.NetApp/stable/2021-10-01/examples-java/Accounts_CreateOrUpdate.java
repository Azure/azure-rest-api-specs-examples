import com.azure.resourcemanager.netapp.models.ActiveDirectory;
import java.util.Arrays;

/** Samples for Accounts CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-10-01/examples/Accounts_CreateOrUpdate.json
     */
    /**
     * Sample code: Accounts_CreateOrUpdate.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void accountsCreateOrUpdate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager
            .accounts()
            .define("account1")
            .withRegion("eastus")
            .withExistingResourceGroup("myRG")
            .withActiveDirectories(
                Arrays
                    .asList(
                        new ActiveDirectory()
                            .withUsername("ad_user_name")
                            .withPassword("ad_password")
                            .withDomain("10.10.10.3")
                            .withDns("10.10.10.3, 10.10.10.4")
                            .withSmbServerName("SMBServer")
                            .withOrganizationalUnit("OU=Engineering")
                            .withSite("SiteName")
                            .withAesEncryption(true)
                            .withLdapSigning(false)
                            .withLdapOverTls(false)))
            .create();
    }
}
