<script lang="ts">
    import { Modal } from "bootstrap";
    export let imgContent: string;

    let shareDom: HTMLElement;
    let shareDialog: Modal;

    $: if (shareDom) {
        if (!shareDialog) {
            shareDialog = new Modal(shareDom);
        }
    }

    export const show = () => {
        shareDialog?.show();
    };

    $: if (shareDialog) {
        if (imgContent) {
            shareDialog.show();
        } else {
            shareDialog.hide();
        }
    }
</script>

<div class="modal" tabindex="-1" bind:this={shareDom}>
    <div class="modal-dialog modal-dialog-centered modal-lg">
        <div class="modal-content">
            <div class="modal-header">
                <h5 class="modal-title">分享</h5>
                <button
                    type="button"
                    class="btn-close"
                    data-bs-dismiss="modal"
                    aria-label="Close"
                ></button>
            </div>
            <div class="modal-body">
                <p>长按图片保存或者分享</p>
                <div
                    class="d-flex flex-column justify-content-center align-items-center"
                >
                    <img src={imgContent} class="w-100" alt="chart" />
                </div>
            </div>
        </div>
    </div>
</div>
