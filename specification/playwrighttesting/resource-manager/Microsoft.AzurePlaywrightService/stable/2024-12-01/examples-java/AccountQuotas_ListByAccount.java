
/**
 * Samples for AccountQuotas ListByAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-12-01/AccountQuotas_ListByAccount.json
     */
    /**
     * Sample code: AccountQuotas_ListByAccount.
     * 
     * @param manager Entry point to PlaywrightTestingManager.
     */
    public static void
        accountQuotasListByAccount(com.azure.resourcemanager.playwrighttesting.PlaywrightTestingManager manager) {
        manager.accountQuotas().listByAccount("dummyrg", "myPlaywrightAccount", com.azure.core.util.Context.NONE);
    }
}
