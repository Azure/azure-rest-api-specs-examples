import com.azure.core.util.Context;

/** Samples for EmailServices ListVerifiedExchangeOnlineDomains. */
public final class Main {
    /*
     * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/preview/2021-10-01-preview/examples/emailServices/getVerifiedExchangeOnlineDomains.json
     */
    /**
     * Sample code: Get verified Exchange Online domains.
     *
     * @param manager Entry point to CommunicationManager.
     */
    public static void getVerifiedExchangeOnlineDomains(
        com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.emailServices().listVerifiedExchangeOnlineDomainsWithResponse(Context.NONE);
    }
}
