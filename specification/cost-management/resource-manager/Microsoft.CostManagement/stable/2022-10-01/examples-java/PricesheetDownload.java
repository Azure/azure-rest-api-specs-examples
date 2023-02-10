/** Samples for PriceSheet Download. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/PricesheetDownload.json
     */
    /**
     * Sample code: PricesheetDownload.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void pricesheetDownload(com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .priceSheets()
            .download(
                "7c05a543-80ff-571e-9f98-1063b3b53cf2:99ad03ad-2d1b-4889-a452-090ad407d25f_2019-05-31",
                "2USN-TPCD-BG7-TGB",
                "T000940677",
                com.azure.core.util.Context.NONE);
    }
}
