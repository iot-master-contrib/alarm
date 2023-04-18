 
import { Component } from '@angular/core';
import { Router } from "@angular/router";
import { RequestService } from "../../request.service";
import { NzMessageService } from "ng-zorro-antd/message";
import { NzTableQueryParams } from "ng-zorro-antd/table";
import { NzModalService } from "ng-zorro-antd/modal"; 
import { tableHeight, onAllChecked, onItemChecked, batchdel, refreshCheckedStatus } from "../../public";
@Component({
  selector: 'app-alarms',
  templateUrl: './alarms.component.html',
  styleUrls: ['./alarms.component.scss']
})
export class AlarmsComponent {
 text!:string
  loading = false
  datum: any[] = []
  total = 1;
  pageSize = 20;
  uploading: Boolean = false;
  pageIndex = 1;
  query: any = {}
  href!: string;
  filterRead = [
    { text: 'true', value: 1 },
    { text: 'false', value: 0 }
  ]
  filterLevel = [
    { text: '1', value: 1 },
    { text: '2', value: 2 },
    { text: '3', value: 3 },
  ]
  checked = false;
  indeterminate = false;
  setOfCheckedId = new Set<number>();
  delResData: any = [];
  constructor(private router: Router,
    private rs: RequestService,
    private modal: NzModalService,
    private msg: NzMessageService
  ) {
    //this.load();
  }
  onSearch(){}
  reload() {
    this.datum = [];
    this.load()
  }

  load() {
    this.loading = true
    this.rs.post("/app/alarm/api/alarm/search", this.query).subscribe(res => {
      const { data, total } = res;
      this.datum = data || [];
      this.total = total || 0;
      this.setOfCheckedId.clear();
      refreshCheckedStatus(this);
    }).add(() => {
      this.loading = false;
    })
  }
 
  delete(id: number, size?: number) {
    this.rs.get(`/app/alarm/api/alarm/${id}/delete`).subscribe(res => {
      if (!size ) {
        this.msg.success("删除成功");
        this.datum = this.datum.filter(d => d.id !== id);
      } else if (size) {
        this.delResData.push(res);
        if (size === this.delResData.length) {
          this.msg.success("删除成功");
          this.load();
        }
      }
    })
  }

  onQuery($event: NzTableQueryParams) {
    // ParseTableQuery($event, this.query)
    // this.load();
  }
  pageIndexChange(pageIndex: number) {
    console.log("pageIndex:", pageIndex)
  }
  pageSizeChange(pageSize: number) {
    this.query.limit = pageSize;
    this.load();
  }

  search($event: string) {
    this.query.keyword = {
      title: $event,
      Message: $event,
    };
    this.query.skip = 0;
    this.load();
  }
  handleImport(e: any) {
    const file: File = e.target.files[0];
    const formData = new FormData();
    formData.append('file', file)
     this.rs.post(`/app/alarm/api/job/import`,formData).subscribe((res)=>{console.log(res )})
  }
  handleExport(){ 
    this.href = `/app/alarm/api/job/export`;
  }
  read(data: any) {
    this.rs.get(`/app/alarm/api/alarm/${data.id}/read`).subscribe(res => {
      data.read = true;
      
    })
  }

  getTableHeight() {
    return tableHeight(this);
  }
  handleBatchDel() {
    batchdel(this);
  }
  handleAllChecked(id: any) {
    onAllChecked(id, this);
  }
  handleItemChecked(id: number, checked: boolean) {
    onItemChecked(id, checked, this);
  }
}
