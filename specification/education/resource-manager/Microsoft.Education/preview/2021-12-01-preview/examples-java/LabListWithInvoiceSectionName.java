import com.azure.core.util.Context;

/** Samples for Labs List. */
public final class Main {
    /*
     * x-ms-original-file: specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/LabListWithInvoiceSectionName.json
     */
    /**
     * Sample code: LabListWithInvoiceSectionName.
     *
     * @param manager Entry point to EducationManager.
     */
    public static void labListWithInvoiceSectionName(com.azure.resourcemanager.education.EducationManager manager) {
        manager.labs().list("{billingAccountName}", "{billingProfileName}", "{invoiceSectionName}", true, Context.NONE);
    }
}
