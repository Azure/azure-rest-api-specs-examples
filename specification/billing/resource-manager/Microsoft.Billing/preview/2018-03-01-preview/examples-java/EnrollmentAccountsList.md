```java
import com.azure.core.util.Context;

/** Samples for EnrollmentAccounts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/preview/2018-03-01-preview/examples/EnrollmentAccountsList.json
     */
    /**
     * Sample code: EnrollmentAccountsList.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void enrollmentAccountsList(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.enrollmentAccounts().list(Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-billing_1.0.0-beta.2/sdk/billing/azure-resourcemanager-billing/README.md) on how to add the SDK to your project and authenticate.
