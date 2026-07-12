
/**
 * Samples for Accounts RefreshLdapBindPassword.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-15-preview/Accounts_RefreshLdapBindPassword.json
     */
    /**
     * Sample code: NetAppAccounts_RefreshLdapBindPassword.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void
        netAppAccountsRefreshLdapBindPassword(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.accounts().refreshLdapBindPassword("myRG", "account1", com.azure.core.util.Context.NONE);
    }
}
