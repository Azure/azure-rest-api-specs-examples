import com.azure.resourcemanager.cognitiveservices.models.Account;
import com.azure.resourcemanager.cognitiveservices.models.Sku;

/** Samples for Accounts Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2023-05-01/examples/UpdateAccount.json
     */
    /**
     * Sample code: Update Account.
     *
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void updateAccount(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        Account resource =
            manager
                .accounts()
                .getByResourceGroupWithResponse("bvttest", "bingSearch", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().withSku(new Sku().withName("S2")).apply();
    }
}
