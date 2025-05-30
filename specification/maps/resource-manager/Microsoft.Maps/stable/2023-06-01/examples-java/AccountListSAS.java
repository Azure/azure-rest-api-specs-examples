
import com.azure.resourcemanager.maps.models.AccountSasParameters;
import com.azure.resourcemanager.maps.models.SigningKey;
import java.util.Arrays;

/**
 * Samples for Accounts ListSas.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/maps/resource-manager/Microsoft.Maps/stable/2023-06-01/examples/AccountListSAS.json
     */
    /**
     * Sample code: List Account Sas.
     * 
     * @param manager Entry point to AzureMapsManager.
     */
    public static void listAccountSas(com.azure.resourcemanager.maps.AzureMapsManager manager) {
        manager.accounts().listSasWithResponse("myResourceGroup", "myMapsAccount",
            new AccountSasParameters().withSigningKey(SigningKey.PRIMARY_KEY)
                .withPrincipalId("e917f87b-324d-4728-98ed-e31d311a7d65").withRegions(Arrays.asList("eastus"))
                .withMaxRatePerSecond(500).withStart("2017-05-24T10:42:03.1567373Z")
                .withExpiry("2017-05-24T11:42:03.1567373Z"),
            com.azure.core.util.Context.NONE);
    }
}
