
import com.azure.resourcemanager.netapp.models.LdapConfiguration;
import java.util.Arrays;

/**
 * Samples for Accounts CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-15-preview/Accounts_CreateOrUpdateLdapConfig.json
     */
    /**
     * Sample code: Accounts_CreateOrUpdateLdapConfig.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void accountsCreateOrUpdateLdapConfig(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.accounts().define("account1").withRegion("eastus").withExistingResourceGroup("myRG")
            .withLdapConfiguration(new LdapConfiguration().withDomain("example.com")
                .withLdapServers(Arrays.asList("192.0.2.1", "192.0.2.2")).withLdapOverTLS(false))
            .create();
    }
}
