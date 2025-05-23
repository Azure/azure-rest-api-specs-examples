
import com.azure.resourcemanager.graphservices.models.AccountResourceProperties;

/**
 * Samples for Accounts CreateAndUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/graphservicesprod/resource-manager/Microsoft.GraphServices/stable/2023-04-13/examples/
     * Accounts_Create.json
     */
    /**
     * Sample code: Create Account resource.
     * 
     * @param manager Entry point to GraphServicesManager.
     */
    public static void createAccountResource(com.azure.resourcemanager.graphservices.GraphServicesManager manager) {
        manager.accounts().define("11111111-aaaa-1111-bbbb-1111111111111").withRegion((String) null)
            .withExistingResourceGroup("testResourceGroupGRAM")
            .withProperties(new AccountResourceProperties().withAppId("11111111-aaaa-1111-bbbb-111111111111")).create();
    }
}
