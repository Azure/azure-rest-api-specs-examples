import com.azure.core.util.Context;

/** Samples for Instructions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/Instruction.json
     */
    /**
     * Sample code: Instruction.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void instruction(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .instructions()
            .getWithResponse("{billingAccountName}", "{billingProfileName}", "{instructionName}", Context.NONE);
    }
}
